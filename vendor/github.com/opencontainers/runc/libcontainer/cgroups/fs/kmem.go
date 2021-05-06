// +build linux,!nokmem

package fs

import ()

func EnableKernelMemoryAccounting(path string) error {
	// Ensure that kernel memory is available in this kernel build. If it
	// isn't, we just ignore it because EnableKernelMemoryAccounting is
	// automatically called for all memory limits.
	// if !cgroups.PathExists(filepath.Join(path, cgroupKernelMemoryLimit)) {
	// 	return nil
	// }
	// We have to limit the kernel memory here as it won't be accounted at all
	// until a limit is set on the cgroup and limit cannot be set once the
	// cgroup has children, or if there are already tasks in the cgroup.
	// for _, i := range []int64{1, -1} {
	// 	if err := setKernelMemory(path, i); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func setKernelMemory(path string, kernelMemoryLimit int64) error {
	return nil
}
