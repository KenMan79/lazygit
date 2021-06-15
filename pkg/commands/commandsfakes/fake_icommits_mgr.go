// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/jesseduffield/lazygit/pkg/commands"
	"github.com/jesseduffield/lazygit/pkg/commands/types"
)

type FakeICommitsMgr struct {
	AmendHeadStub        func() error
	amendHeadMutex       sync.RWMutex
	amendHeadArgsForCall []struct {
	}
	amendHeadReturns struct {
		result1 error
	}
	amendHeadReturnsOnCall map[int]struct {
		result1 error
	}
	AmendHeadCmdObjStub        func() types.ICmdObj
	amendHeadCmdObjMutex       sync.RWMutex
	amendHeadCmdObjArgsForCall []struct {
	}
	amendHeadCmdObjReturns struct {
		result1 types.ICmdObj
	}
	amendHeadCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	CommitCmdObjStub        func(string, string) types.ICmdObj
	commitCmdObjMutex       sync.RWMutex
	commitCmdObjArgsForCall []struct {
		arg1 string
		arg2 string
	}
	commitCmdObjReturns struct {
		result1 types.ICmdObj
	}
	commitCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	CreateFixupCommitStub        func(string) error
	createFixupCommitMutex       sync.RWMutex
	createFixupCommitArgsForCall []struct {
		arg1 string
	}
	createFixupCommitReturns struct {
		result1 error
	}
	createFixupCommitReturnsOnCall map[int]struct {
		result1 error
	}
	GetHeadMessageStub        func() (string, error)
	getHeadMessageMutex       sync.RWMutex
	getHeadMessageArgsForCall []struct {
	}
	getHeadMessageReturns struct {
		result1 string
		result2 error
	}
	getHeadMessageReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetMessageStub        func(string) (string, error)
	getMessageMutex       sync.RWMutex
	getMessageArgsForCall []struct {
		arg1 string
	}
	getMessageReturns struct {
		result1 string
		result2 error
	}
	getMessageReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetMessageFirstLineStub        func(string) (string, error)
	getMessageFirstLineMutex       sync.RWMutex
	getMessageFirstLineArgsForCall []struct {
		arg1 string
	}
	getMessageFirstLineReturns struct {
		result1 string
		result2 error
	}
	getMessageFirstLineReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RevertStub        func(string) error
	revertMutex       sync.RWMutex
	revertArgsForCall []struct {
		arg1 string
	}
	revertReturns struct {
		result1 error
	}
	revertReturnsOnCall map[int]struct {
		result1 error
	}
	RevertMergeStub        func(string, int) error
	revertMergeMutex       sync.RWMutex
	revertMergeArgsForCall []struct {
		arg1 string
		arg2 int
	}
	revertMergeReturns struct {
		result1 error
	}
	revertMergeReturnsOnCall map[int]struct {
		result1 error
	}
	RewordHeadStub        func(string) error
	rewordHeadMutex       sync.RWMutex
	rewordHeadArgsForCall []struct {
		arg1 string
	}
	rewordHeadReturns struct {
		result1 error
	}
	rewordHeadReturnsOnCall map[int]struct {
		result1 error
	}
	ShowCmdObjStub        func(string, string) types.ICmdObj
	showCmdObjMutex       sync.RWMutex
	showCmdObjArgsForCall []struct {
		arg1 string
		arg2 string
	}
	showCmdObjReturns struct {
		result1 types.ICmdObj
	}
	showCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeICommitsMgr) AmendHead() error {
	fake.amendHeadMutex.Lock()
	ret, specificReturn := fake.amendHeadReturnsOnCall[len(fake.amendHeadArgsForCall)]
	fake.amendHeadArgsForCall = append(fake.amendHeadArgsForCall, struct {
	}{})
	stub := fake.AmendHeadStub
	fakeReturns := fake.amendHeadReturns
	fake.recordInvocation("AmendHead", []interface{}{})
	fake.amendHeadMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) AmendHeadCallCount() int {
	fake.amendHeadMutex.RLock()
	defer fake.amendHeadMutex.RUnlock()
	return len(fake.amendHeadArgsForCall)
}

func (fake *FakeICommitsMgr) AmendHeadCalls(stub func() error) {
	fake.amendHeadMutex.Lock()
	defer fake.amendHeadMutex.Unlock()
	fake.AmendHeadStub = stub
}

func (fake *FakeICommitsMgr) AmendHeadReturns(result1 error) {
	fake.amendHeadMutex.Lock()
	defer fake.amendHeadMutex.Unlock()
	fake.AmendHeadStub = nil
	fake.amendHeadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) AmendHeadReturnsOnCall(i int, result1 error) {
	fake.amendHeadMutex.Lock()
	defer fake.amendHeadMutex.Unlock()
	fake.AmendHeadStub = nil
	if fake.amendHeadReturnsOnCall == nil {
		fake.amendHeadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.amendHeadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) AmendHeadCmdObj() types.ICmdObj {
	fake.amendHeadCmdObjMutex.Lock()
	ret, specificReturn := fake.amendHeadCmdObjReturnsOnCall[len(fake.amendHeadCmdObjArgsForCall)]
	fake.amendHeadCmdObjArgsForCall = append(fake.amendHeadCmdObjArgsForCall, struct {
	}{})
	stub := fake.AmendHeadCmdObjStub
	fakeReturns := fake.amendHeadCmdObjReturns
	fake.recordInvocation("AmendHeadCmdObj", []interface{}{})
	fake.amendHeadCmdObjMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) AmendHeadCmdObjCallCount() int {
	fake.amendHeadCmdObjMutex.RLock()
	defer fake.amendHeadCmdObjMutex.RUnlock()
	return len(fake.amendHeadCmdObjArgsForCall)
}

func (fake *FakeICommitsMgr) AmendHeadCmdObjCalls(stub func() types.ICmdObj) {
	fake.amendHeadCmdObjMutex.Lock()
	defer fake.amendHeadCmdObjMutex.Unlock()
	fake.AmendHeadCmdObjStub = stub
}

func (fake *FakeICommitsMgr) AmendHeadCmdObjReturns(result1 types.ICmdObj) {
	fake.amendHeadCmdObjMutex.Lock()
	defer fake.amendHeadCmdObjMutex.Unlock()
	fake.AmendHeadCmdObjStub = nil
	fake.amendHeadCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) AmendHeadCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.amendHeadCmdObjMutex.Lock()
	defer fake.amendHeadCmdObjMutex.Unlock()
	fake.AmendHeadCmdObjStub = nil
	if fake.amendHeadCmdObjReturnsOnCall == nil {
		fake.amendHeadCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.amendHeadCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) CommitCmdObj(arg1 string, arg2 string) types.ICmdObj {
	fake.commitCmdObjMutex.Lock()
	ret, specificReturn := fake.commitCmdObjReturnsOnCall[len(fake.commitCmdObjArgsForCall)]
	fake.commitCmdObjArgsForCall = append(fake.commitCmdObjArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CommitCmdObjStub
	fakeReturns := fake.commitCmdObjReturns
	fake.recordInvocation("CommitCmdObj", []interface{}{arg1, arg2})
	fake.commitCmdObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) CommitCmdObjCallCount() int {
	fake.commitCmdObjMutex.RLock()
	defer fake.commitCmdObjMutex.RUnlock()
	return len(fake.commitCmdObjArgsForCall)
}

func (fake *FakeICommitsMgr) CommitCmdObjCalls(stub func(string, string) types.ICmdObj) {
	fake.commitCmdObjMutex.Lock()
	defer fake.commitCmdObjMutex.Unlock()
	fake.CommitCmdObjStub = stub
}

func (fake *FakeICommitsMgr) CommitCmdObjArgsForCall(i int) (string, string) {
	fake.commitCmdObjMutex.RLock()
	defer fake.commitCmdObjMutex.RUnlock()
	argsForCall := fake.commitCmdObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeICommitsMgr) CommitCmdObjReturns(result1 types.ICmdObj) {
	fake.commitCmdObjMutex.Lock()
	defer fake.commitCmdObjMutex.Unlock()
	fake.CommitCmdObjStub = nil
	fake.commitCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) CommitCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.commitCmdObjMutex.Lock()
	defer fake.commitCmdObjMutex.Unlock()
	fake.CommitCmdObjStub = nil
	if fake.commitCmdObjReturnsOnCall == nil {
		fake.commitCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.commitCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) CreateFixupCommit(arg1 string) error {
	fake.createFixupCommitMutex.Lock()
	ret, specificReturn := fake.createFixupCommitReturnsOnCall[len(fake.createFixupCommitArgsForCall)]
	fake.createFixupCommitArgsForCall = append(fake.createFixupCommitArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateFixupCommitStub
	fakeReturns := fake.createFixupCommitReturns
	fake.recordInvocation("CreateFixupCommit", []interface{}{arg1})
	fake.createFixupCommitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) CreateFixupCommitCallCount() int {
	fake.createFixupCommitMutex.RLock()
	defer fake.createFixupCommitMutex.RUnlock()
	return len(fake.createFixupCommitArgsForCall)
}

func (fake *FakeICommitsMgr) CreateFixupCommitCalls(stub func(string) error) {
	fake.createFixupCommitMutex.Lock()
	defer fake.createFixupCommitMutex.Unlock()
	fake.CreateFixupCommitStub = stub
}

func (fake *FakeICommitsMgr) CreateFixupCommitArgsForCall(i int) string {
	fake.createFixupCommitMutex.RLock()
	defer fake.createFixupCommitMutex.RUnlock()
	argsForCall := fake.createFixupCommitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICommitsMgr) CreateFixupCommitReturns(result1 error) {
	fake.createFixupCommitMutex.Lock()
	defer fake.createFixupCommitMutex.Unlock()
	fake.CreateFixupCommitStub = nil
	fake.createFixupCommitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) CreateFixupCommitReturnsOnCall(i int, result1 error) {
	fake.createFixupCommitMutex.Lock()
	defer fake.createFixupCommitMutex.Unlock()
	fake.CreateFixupCommitStub = nil
	if fake.createFixupCommitReturnsOnCall == nil {
		fake.createFixupCommitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createFixupCommitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) GetHeadMessage() (string, error) {
	fake.getHeadMessageMutex.Lock()
	ret, specificReturn := fake.getHeadMessageReturnsOnCall[len(fake.getHeadMessageArgsForCall)]
	fake.getHeadMessageArgsForCall = append(fake.getHeadMessageArgsForCall, struct {
	}{})
	stub := fake.GetHeadMessageStub
	fakeReturns := fake.getHeadMessageReturns
	fake.recordInvocation("GetHeadMessage", []interface{}{})
	fake.getHeadMessageMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeICommitsMgr) GetHeadMessageCallCount() int {
	fake.getHeadMessageMutex.RLock()
	defer fake.getHeadMessageMutex.RUnlock()
	return len(fake.getHeadMessageArgsForCall)
}

func (fake *FakeICommitsMgr) GetHeadMessageCalls(stub func() (string, error)) {
	fake.getHeadMessageMutex.Lock()
	defer fake.getHeadMessageMutex.Unlock()
	fake.GetHeadMessageStub = stub
}

func (fake *FakeICommitsMgr) GetHeadMessageReturns(result1 string, result2 error) {
	fake.getHeadMessageMutex.Lock()
	defer fake.getHeadMessageMutex.Unlock()
	fake.GetHeadMessageStub = nil
	fake.getHeadMessageReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) GetHeadMessageReturnsOnCall(i int, result1 string, result2 error) {
	fake.getHeadMessageMutex.Lock()
	defer fake.getHeadMessageMutex.Unlock()
	fake.GetHeadMessageStub = nil
	if fake.getHeadMessageReturnsOnCall == nil {
		fake.getHeadMessageReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getHeadMessageReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) GetMessage(arg1 string) (string, error) {
	fake.getMessageMutex.Lock()
	ret, specificReturn := fake.getMessageReturnsOnCall[len(fake.getMessageArgsForCall)]
	fake.getMessageArgsForCall = append(fake.getMessageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetMessageStub
	fakeReturns := fake.getMessageReturns
	fake.recordInvocation("GetMessage", []interface{}{arg1})
	fake.getMessageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeICommitsMgr) GetMessageCallCount() int {
	fake.getMessageMutex.RLock()
	defer fake.getMessageMutex.RUnlock()
	return len(fake.getMessageArgsForCall)
}

func (fake *FakeICommitsMgr) GetMessageCalls(stub func(string) (string, error)) {
	fake.getMessageMutex.Lock()
	defer fake.getMessageMutex.Unlock()
	fake.GetMessageStub = stub
}

func (fake *FakeICommitsMgr) GetMessageArgsForCall(i int) string {
	fake.getMessageMutex.RLock()
	defer fake.getMessageMutex.RUnlock()
	argsForCall := fake.getMessageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICommitsMgr) GetMessageReturns(result1 string, result2 error) {
	fake.getMessageMutex.Lock()
	defer fake.getMessageMutex.Unlock()
	fake.GetMessageStub = nil
	fake.getMessageReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) GetMessageReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMessageMutex.Lock()
	defer fake.getMessageMutex.Unlock()
	fake.GetMessageStub = nil
	if fake.getMessageReturnsOnCall == nil {
		fake.getMessageReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getMessageReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) GetMessageFirstLine(arg1 string) (string, error) {
	fake.getMessageFirstLineMutex.Lock()
	ret, specificReturn := fake.getMessageFirstLineReturnsOnCall[len(fake.getMessageFirstLineArgsForCall)]
	fake.getMessageFirstLineArgsForCall = append(fake.getMessageFirstLineArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetMessageFirstLineStub
	fakeReturns := fake.getMessageFirstLineReturns
	fake.recordInvocation("GetMessageFirstLine", []interface{}{arg1})
	fake.getMessageFirstLineMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeICommitsMgr) GetMessageFirstLineCallCount() int {
	fake.getMessageFirstLineMutex.RLock()
	defer fake.getMessageFirstLineMutex.RUnlock()
	return len(fake.getMessageFirstLineArgsForCall)
}

func (fake *FakeICommitsMgr) GetMessageFirstLineCalls(stub func(string) (string, error)) {
	fake.getMessageFirstLineMutex.Lock()
	defer fake.getMessageFirstLineMutex.Unlock()
	fake.GetMessageFirstLineStub = stub
}

func (fake *FakeICommitsMgr) GetMessageFirstLineArgsForCall(i int) string {
	fake.getMessageFirstLineMutex.RLock()
	defer fake.getMessageFirstLineMutex.RUnlock()
	argsForCall := fake.getMessageFirstLineArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICommitsMgr) GetMessageFirstLineReturns(result1 string, result2 error) {
	fake.getMessageFirstLineMutex.Lock()
	defer fake.getMessageFirstLineMutex.Unlock()
	fake.GetMessageFirstLineStub = nil
	fake.getMessageFirstLineReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) GetMessageFirstLineReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMessageFirstLineMutex.Lock()
	defer fake.getMessageFirstLineMutex.Unlock()
	fake.GetMessageFirstLineStub = nil
	if fake.getMessageFirstLineReturnsOnCall == nil {
		fake.getMessageFirstLineReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getMessageFirstLineReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeICommitsMgr) Revert(arg1 string) error {
	fake.revertMutex.Lock()
	ret, specificReturn := fake.revertReturnsOnCall[len(fake.revertArgsForCall)]
	fake.revertArgsForCall = append(fake.revertArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RevertStub
	fakeReturns := fake.revertReturns
	fake.recordInvocation("Revert", []interface{}{arg1})
	fake.revertMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) RevertCallCount() int {
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	return len(fake.revertArgsForCall)
}

func (fake *FakeICommitsMgr) RevertCalls(stub func(string) error) {
	fake.revertMutex.Lock()
	defer fake.revertMutex.Unlock()
	fake.RevertStub = stub
}

func (fake *FakeICommitsMgr) RevertArgsForCall(i int) string {
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	argsForCall := fake.revertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICommitsMgr) RevertReturns(result1 error) {
	fake.revertMutex.Lock()
	defer fake.revertMutex.Unlock()
	fake.RevertStub = nil
	fake.revertReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) RevertReturnsOnCall(i int, result1 error) {
	fake.revertMutex.Lock()
	defer fake.revertMutex.Unlock()
	fake.RevertStub = nil
	if fake.revertReturnsOnCall == nil {
		fake.revertReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.revertReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) RevertMerge(arg1 string, arg2 int) error {
	fake.revertMergeMutex.Lock()
	ret, specificReturn := fake.revertMergeReturnsOnCall[len(fake.revertMergeArgsForCall)]
	fake.revertMergeArgsForCall = append(fake.revertMergeArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	stub := fake.RevertMergeStub
	fakeReturns := fake.revertMergeReturns
	fake.recordInvocation("RevertMerge", []interface{}{arg1, arg2})
	fake.revertMergeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) RevertMergeCallCount() int {
	fake.revertMergeMutex.RLock()
	defer fake.revertMergeMutex.RUnlock()
	return len(fake.revertMergeArgsForCall)
}

func (fake *FakeICommitsMgr) RevertMergeCalls(stub func(string, int) error) {
	fake.revertMergeMutex.Lock()
	defer fake.revertMergeMutex.Unlock()
	fake.RevertMergeStub = stub
}

func (fake *FakeICommitsMgr) RevertMergeArgsForCall(i int) (string, int) {
	fake.revertMergeMutex.RLock()
	defer fake.revertMergeMutex.RUnlock()
	argsForCall := fake.revertMergeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeICommitsMgr) RevertMergeReturns(result1 error) {
	fake.revertMergeMutex.Lock()
	defer fake.revertMergeMutex.Unlock()
	fake.RevertMergeStub = nil
	fake.revertMergeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) RevertMergeReturnsOnCall(i int, result1 error) {
	fake.revertMergeMutex.Lock()
	defer fake.revertMergeMutex.Unlock()
	fake.RevertMergeStub = nil
	if fake.revertMergeReturnsOnCall == nil {
		fake.revertMergeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.revertMergeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) RewordHead(arg1 string) error {
	fake.rewordHeadMutex.Lock()
	ret, specificReturn := fake.rewordHeadReturnsOnCall[len(fake.rewordHeadArgsForCall)]
	fake.rewordHeadArgsForCall = append(fake.rewordHeadArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RewordHeadStub
	fakeReturns := fake.rewordHeadReturns
	fake.recordInvocation("RewordHead", []interface{}{arg1})
	fake.rewordHeadMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) RewordHeadCallCount() int {
	fake.rewordHeadMutex.RLock()
	defer fake.rewordHeadMutex.RUnlock()
	return len(fake.rewordHeadArgsForCall)
}

func (fake *FakeICommitsMgr) RewordHeadCalls(stub func(string) error) {
	fake.rewordHeadMutex.Lock()
	defer fake.rewordHeadMutex.Unlock()
	fake.RewordHeadStub = stub
}

func (fake *FakeICommitsMgr) RewordHeadArgsForCall(i int) string {
	fake.rewordHeadMutex.RLock()
	defer fake.rewordHeadMutex.RUnlock()
	argsForCall := fake.rewordHeadArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICommitsMgr) RewordHeadReturns(result1 error) {
	fake.rewordHeadMutex.Lock()
	defer fake.rewordHeadMutex.Unlock()
	fake.RewordHeadStub = nil
	fake.rewordHeadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) RewordHeadReturnsOnCall(i int, result1 error) {
	fake.rewordHeadMutex.Lock()
	defer fake.rewordHeadMutex.Unlock()
	fake.RewordHeadStub = nil
	if fake.rewordHeadReturnsOnCall == nil {
		fake.rewordHeadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rewordHeadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICommitsMgr) ShowCmdObj(arg1 string, arg2 string) types.ICmdObj {
	fake.showCmdObjMutex.Lock()
	ret, specificReturn := fake.showCmdObjReturnsOnCall[len(fake.showCmdObjArgsForCall)]
	fake.showCmdObjArgsForCall = append(fake.showCmdObjArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ShowCmdObjStub
	fakeReturns := fake.showCmdObjReturns
	fake.recordInvocation("ShowCmdObj", []interface{}{arg1, arg2})
	fake.showCmdObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICommitsMgr) ShowCmdObjCallCount() int {
	fake.showCmdObjMutex.RLock()
	defer fake.showCmdObjMutex.RUnlock()
	return len(fake.showCmdObjArgsForCall)
}

func (fake *FakeICommitsMgr) ShowCmdObjCalls(stub func(string, string) types.ICmdObj) {
	fake.showCmdObjMutex.Lock()
	defer fake.showCmdObjMutex.Unlock()
	fake.ShowCmdObjStub = stub
}

func (fake *FakeICommitsMgr) ShowCmdObjArgsForCall(i int) (string, string) {
	fake.showCmdObjMutex.RLock()
	defer fake.showCmdObjMutex.RUnlock()
	argsForCall := fake.showCmdObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeICommitsMgr) ShowCmdObjReturns(result1 types.ICmdObj) {
	fake.showCmdObjMutex.Lock()
	defer fake.showCmdObjMutex.Unlock()
	fake.ShowCmdObjStub = nil
	fake.showCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) ShowCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.showCmdObjMutex.Lock()
	defer fake.showCmdObjMutex.Unlock()
	fake.ShowCmdObjStub = nil
	if fake.showCmdObjReturnsOnCall == nil {
		fake.showCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.showCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeICommitsMgr) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.amendHeadMutex.RLock()
	defer fake.amendHeadMutex.RUnlock()
	fake.amendHeadCmdObjMutex.RLock()
	defer fake.amendHeadCmdObjMutex.RUnlock()
	fake.commitCmdObjMutex.RLock()
	defer fake.commitCmdObjMutex.RUnlock()
	fake.createFixupCommitMutex.RLock()
	defer fake.createFixupCommitMutex.RUnlock()
	fake.getHeadMessageMutex.RLock()
	defer fake.getHeadMessageMutex.RUnlock()
	fake.getMessageMutex.RLock()
	defer fake.getMessageMutex.RUnlock()
	fake.getMessageFirstLineMutex.RLock()
	defer fake.getMessageFirstLineMutex.RUnlock()
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	fake.revertMergeMutex.RLock()
	defer fake.revertMergeMutex.RUnlock()
	fake.rewordHeadMutex.RLock()
	defer fake.rewordHeadMutex.RUnlock()
	fake.showCmdObjMutex.RLock()
	defer fake.showCmdObjMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeICommitsMgr) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ commands.ICommitsMgr = new(FakeICommitsMgr)
