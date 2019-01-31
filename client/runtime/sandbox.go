package runtime

import "log"

// TODO: implement based on rust implementation

// Sandbox ...
type Sandbox struct {
}

// NewSandbox ...
func NewSandbox(env Env) *Sandbox {
	return &Sandbox{}
}

// Instantiate ...
func (s *Sandbox) Instantiate(a, b, c, d, e, f int64) int64 {
	log.Fatal("sandbox_instantiate not implemented")
	return 0
}

// InstanceTeardown ...
func (s *Sandbox) InstanceTeardown(instanceIndex int64) {
	log.Fatal("sandbox_instance_teardown not implemented")
}

// Invoke ...
func (s *Sandbox) Invoke(instanceIndex int64, exportPtr Pointer, exportLength int64, argsPtr Pointer, argsLength int64, returnValPtr Pointer, returnValLength int64, state int64) {
	log.Fatal("sandbox_invoke not implemented")
}

// MemoryGet ...
func (s *Sandbox) MemoryGet(memoryIndex int64, offset int64, ptr Pointer, length int64) int64 {
	log.Fatal("sandbox_memory_get not implemented")
	return 0
}

// MemoryNew ...
func (s *Sandbox) MemoryNew(initial int64, maximum int64) int64 {
	log.Fatal("sandbox_memory_new not implemented")
	return 0
}

// MemorySet ...
func (s *Sandbox) MemorySet(memoryIndex int64, offset int64, ptr Pointer, length int64) int64 {
	log.Fatal("sandbox_memory_set not implemented")
	return 0
}

// MemoryTeardown ...
func (s *Sandbox) MemoryTeardown(memoryIndex int64) {
	log.Fatal("sandbox_memory_teardown not implemented")
}
