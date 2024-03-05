package spark

import (
	"context"
	"github.com/shenfz/TabakoAssistant/backend/global"
	"sort"
	"strings"
)

/**
 * @Author shenfz
 * @Date 2024/3/4 15:25
 * @Email 1328919715@qq.com
 * @Description:
 **/

func (s *SparkWsConn) SendMsg(ctx context.Context, obj SparkReqMsg) string {
	select {
	case s.sendChan <- obj:
		return "sent to chan "
	case <-ctx.Done():
		return "reply timeout[5s]"
	}
	return "unknown"
}

func (s *SparkWsConn) CheckPayloadIsExisted(key string) bool {
	s.locker.RLocker()
	defer s.locker.RUnlock()
	_, ok := s.payloadCache[key]
	return ok
}

func (s *SparkWsConn) CreatePayload(key string, sentText []Text) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.payloadCache[key] = sentText
}

// UpUserPayload record sent text as payload
func (s *SparkWsConn) UpUserPayload(sentText []Text) []Text {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.payloadCache[global.SparkAPI_Role_User] = append(s.payloadCache[global.SparkAPI_Role_User], sentText...)
	return s.payloadCache[global.SparkAPI_Role_User]
}

// ClearPayload clear
func (s *SparkWsConn) ClearPayload(key string) {
	s.locker.Lock()
	defer s.locker.Unlock()
	delete(s.payloadCache, key)
	s.logger.Debugf("clear payload by KEY = [%s] ", key)
}

func (s *SparkWsConn) GetPayload(key string) []Text {
	s.locker.RLocker()
	defer s.locker.RUnlock()
	dst, ok := s.payloadCache[key]
	if !ok {
		s.logger.Warnf("no any payload KEY = [%s]", key)
	}
	return dst
}

func (s *SparkWsConn) AssembleStreamContent() string {

	var dst = &strings.Builder{}
	defer dst.Reset()
	var sortedKeys []string

	s.locker.RLock()
	for key, _ := range s.payloadCache {
		if strings.Contains(key, global.SparkAPI_Role_Assistant) {
			sortedKeys = append(sortedKeys, key)
		}
	}
	s.locker.RUnlock()

	sort.Slice(sortedKeys, func(i, j int) bool {
		numI := []byte(strings.Split(sortedKeys[i], "-")[0])
		numJ := []byte(strings.Split(sortedKeys[j], "-")[0])
		if numI[0] > numJ[0] {
			return false
		}
		return true
	})

	s.locker.Lock()
	for _, key := range sortedKeys {
		temp := s.payloadCache[key]
		// add to user payload
		s.payloadCache[global.SparkAPI_Role_User] = append(s.payloadCache[global.SparkAPI_Role_User], temp...)
		// write
		for _, text := range temp {
			dst.WriteString(text.Content)
		}
		// delete
		delete(s.payloadCache, key)
	}
	s.locker.Unlock()

	//	s.logger.Infof("Get Assistant Assemble Content: %s", dst.String())
	//
	return dst.String()
}

func (s *SparkWsConn) SaveAssistantContentBySeq(key string, texts []Text) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.payloadCache[key] = append(s.payloadCache[key], texts...)
}
