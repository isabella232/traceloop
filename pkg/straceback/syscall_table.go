// generated by ./scripts/get_syscall_table.sh
package straceback

func init() {
	syscallNames = make(map[int]string)
	syscallNames[0] = "read"
	syscallNames[1] = "write"
	syscallNames[2] = "open"
	syscallNames[3] = "close"
	syscallNames[4] = "stat"
	syscallNames[5] = "fstat"
	syscallNames[6] = "lstat"
	syscallNames[7] = "poll"
	syscallNames[8] = "lseek"
	syscallNames[9] = "mmap"
	syscallNames[10] = "mprotect"
	syscallNames[11] = "munmap"
	syscallNames[12] = "brk"
	syscallNames[13] = "rt_sigaction"
	syscallNames[14] = "rt_sigprocmask"
	syscallNames[15] = "rt_sigreturn"
	syscallNames[16] = "ioctl"
	syscallNames[17] = "pread64"
	syscallNames[18] = "pwrite64"
	syscallNames[19] = "readv"
	syscallNames[20] = "writev"
	syscallNames[21] = "access"
	syscallNames[22] = "pipe"
	syscallNames[23] = "select"
	syscallNames[24] = "sched_yield"
	syscallNames[25] = "mremap"
	syscallNames[26] = "msync"
	syscallNames[27] = "mincore"
	syscallNames[28] = "madvise"
	syscallNames[29] = "shmget"
	syscallNames[30] = "shmat"
	syscallNames[31] = "shmctl"
	syscallNames[32] = "dup"
	syscallNames[33] = "dup2"
	syscallNames[34] = "pause"
	syscallNames[35] = "nanosleep"
	syscallNames[36] = "getitimer"
	syscallNames[37] = "alarm"
	syscallNames[38] = "setitimer"
	syscallNames[39] = "getpid"
	syscallNames[40] = "sendfile"
	syscallNames[41] = "socket"
	syscallNames[42] = "connect"
	syscallNames[43] = "accept"
	syscallNames[44] = "sendto"
	syscallNames[45] = "recvfrom"
	syscallNames[46] = "sendmsg"
	syscallNames[47] = "recvmsg"
	syscallNames[48] = "shutdown"
	syscallNames[49] = "bind"
	syscallNames[50] = "listen"
	syscallNames[51] = "getsockname"
	syscallNames[52] = "getpeername"
	syscallNames[53] = "socketpair"
	syscallNames[54] = "setsockopt"
	syscallNames[55] = "getsockopt"
	syscallNames[56] = "clone"
	syscallNames[57] = "fork"
	syscallNames[58] = "vfork"
	syscallNames[59] = "execve"
	syscallNames[60] = "exit"
	syscallNames[61] = "wait4"
	syscallNames[62] = "kill"
	syscallNames[63] = "uname"
	syscallNames[64] = "semget"
	syscallNames[65] = "semop"
	syscallNames[66] = "semctl"
	syscallNames[67] = "shmdt"
	syscallNames[68] = "msgget"
	syscallNames[69] = "msgsnd"
	syscallNames[70] = "msgrcv"
	syscallNames[71] = "msgctl"
	syscallNames[72] = "fcntl"
	syscallNames[73] = "flock"
	syscallNames[74] = "fsync"
	syscallNames[75] = "fdatasync"
	syscallNames[76] = "truncate"
	syscallNames[77] = "ftruncate"
	syscallNames[78] = "getdents"
	syscallNames[79] = "getcwd"
	syscallNames[80] = "chdir"
	syscallNames[81] = "fchdir"
	syscallNames[82] = "rename"
	syscallNames[83] = "mkdir"
	syscallNames[84] = "rmdir"
	syscallNames[85] = "creat"
	syscallNames[86] = "link"
	syscallNames[87] = "unlink"
	syscallNames[88] = "symlink"
	syscallNames[89] = "readlink"
	syscallNames[90] = "chmod"
	syscallNames[91] = "fchmod"
	syscallNames[92] = "chown"
	syscallNames[93] = "fchown"
	syscallNames[94] = "lchown"
	syscallNames[95] = "umask"
	syscallNames[96] = "gettimeofday"
	syscallNames[97] = "getrlimit"
	syscallNames[98] = "getrusage"
	syscallNames[99] = "sysinfo"
	syscallNames[100] = "times"
	syscallNames[101] = "ptrace"
	syscallNames[102] = "getuid"
	syscallNames[103] = "syslog"
	syscallNames[104] = "getgid"
	syscallNames[105] = "setuid"
	syscallNames[106] = "setgid"
	syscallNames[107] = "geteuid"
	syscallNames[108] = "getegid"
	syscallNames[109] = "setpgid"
	syscallNames[110] = "getppid"
	syscallNames[111] = "getpgrp"
	syscallNames[112] = "setsid"
	syscallNames[113] = "setreuid"
	syscallNames[114] = "setregid"
	syscallNames[115] = "getgroups"
	syscallNames[116] = "setgroups"
	syscallNames[117] = "setresuid"
	syscallNames[118] = "getresuid"
	syscallNames[119] = "setresgid"
	syscallNames[120] = "getresgid"
	syscallNames[121] = "getpgid"
	syscallNames[122] = "setfsuid"
	syscallNames[123] = "setfsgid"
	syscallNames[124] = "getsid"
	syscallNames[125] = "capget"
	syscallNames[126] = "capset"
	syscallNames[127] = "rt_sigpending"
	syscallNames[128] = "rt_sigtimedwait"
	syscallNames[129] = "rt_sigqueueinfo"
	syscallNames[130] = "rt_sigsuspend"
	syscallNames[131] = "sigaltstack"
	syscallNames[132] = "utime"
	syscallNames[133] = "mknod"
	syscallNames[134] = "uselib"
	syscallNames[135] = "personality"
	syscallNames[136] = "ustat"
	syscallNames[137] = "statfs"
	syscallNames[138] = "fstatfs"
	syscallNames[139] = "sysfs"
	syscallNames[140] = "getpriority"
	syscallNames[141] = "setpriority"
	syscallNames[142] = "sched_setparam"
	syscallNames[143] = "sched_getparam"
	syscallNames[144] = "sched_setscheduler"
	syscallNames[145] = "sched_getscheduler"
	syscallNames[146] = "sched_get_priority_max"
	syscallNames[147] = "sched_get_priority_min"
	syscallNames[148] = "sched_rr_get_interval"
	syscallNames[149] = "mlock"
	syscallNames[150] = "munlock"
	syscallNames[151] = "mlockall"
	syscallNames[152] = "munlockall"
	syscallNames[153] = "vhangup"
	syscallNames[154] = "modify_ldt"
	syscallNames[155] = "pivot_root"
	syscallNames[156] = "_sysctl"
	syscallNames[157] = "prctl"
	syscallNames[158] = "arch_prctl"
	syscallNames[159] = "adjtimex"
	syscallNames[160] = "setrlimit"
	syscallNames[161] = "chroot"
	syscallNames[162] = "sync"
	syscallNames[163] = "acct"
	syscallNames[164] = "settimeofday"
	syscallNames[165] = "mount"
	syscallNames[166] = "umount2"
	syscallNames[167] = "swapon"
	syscallNames[168] = "swapoff"
	syscallNames[169] = "reboot"
	syscallNames[170] = "sethostname"
	syscallNames[171] = "setdomainname"
	syscallNames[172] = "iopl"
	syscallNames[173] = "ioperm"
	syscallNames[174] = "create_module"
	syscallNames[175] = "init_module"
	syscallNames[176] = "delete_module"
	syscallNames[177] = "get_kernel_syms"
	syscallNames[178] = "query_module"
	syscallNames[179] = "quotactl"
	syscallNames[180] = "nfsservctl"
	syscallNames[181] = "getpmsg"
	syscallNames[182] = "putpmsg"
	syscallNames[183] = "afs_syscall"
	syscallNames[184] = "tuxcall"
	syscallNames[185] = "security"
	syscallNames[186] = "gettid"
	syscallNames[187] = "readahead"
	syscallNames[188] = "setxattr"
	syscallNames[189] = "lsetxattr"
	syscallNames[190] = "fsetxattr"
	syscallNames[191] = "getxattr"
	syscallNames[192] = "lgetxattr"
	syscallNames[193] = "fgetxattr"
	syscallNames[194] = "listxattr"
	syscallNames[195] = "llistxattr"
	syscallNames[196] = "flistxattr"
	syscallNames[197] = "removexattr"
	syscallNames[198] = "lremovexattr"
	syscallNames[199] = "fremovexattr"
	syscallNames[200] = "tkill"
	syscallNames[201] = "time"
	syscallNames[202] = "futex"
	syscallNames[203] = "sched_setaffinity"
	syscallNames[204] = "sched_getaffinity"
	syscallNames[205] = "set_thread_area"
	syscallNames[206] = "io_setup"
	syscallNames[207] = "io_destroy"
	syscallNames[208] = "io_getevents"
	syscallNames[209] = "io_submit"
	syscallNames[210] = "io_cancel"
	syscallNames[211] = "get_thread_area"
	syscallNames[212] = "lookup_dcookie"
	syscallNames[213] = "epoll_create"
	syscallNames[214] = "epoll_ctl_old"
	syscallNames[215] = "epoll_wait_old"
	syscallNames[216] = "remap_file_pages"
	syscallNames[217] = "getdents64"
	syscallNames[218] = "set_tid_address"
	syscallNames[219] = "restart_syscall"
	syscallNames[220] = "semtimedop"
	syscallNames[221] = "fadvise64"
	syscallNames[222] = "timer_create"
	syscallNames[223] = "timer_settime"
	syscallNames[224] = "timer_gettime"
	syscallNames[225] = "timer_getoverrun"
	syscallNames[226] = "timer_delete"
	syscallNames[227] = "clock_settime"
	syscallNames[228] = "clock_gettime"
	syscallNames[229] = "clock_getres"
	syscallNames[230] = "clock_nanosleep"
	syscallNames[231] = "exit_group"
	syscallNames[232] = "epoll_wait"
	syscallNames[233] = "epoll_ctl"
	syscallNames[234] = "tgkill"
	syscallNames[235] = "utimes"
	syscallNames[236] = "vserver"
	syscallNames[237] = "mbind"
	syscallNames[238] = "set_mempolicy"
	syscallNames[239] = "get_mempolicy"
	syscallNames[240] = "mq_open"
	syscallNames[241] = "mq_unlink"
	syscallNames[242] = "mq_timedsend"
	syscallNames[243] = "mq_timedreceive"
	syscallNames[244] = "mq_notify"
	syscallNames[245] = "mq_getsetattr"
	syscallNames[246] = "kexec_load"
	syscallNames[247] = "waitid"
	syscallNames[248] = "add_key"
	syscallNames[249] = "request_key"
	syscallNames[250] = "keyctl"
	syscallNames[251] = "ioprio_set"
	syscallNames[252] = "ioprio_get"
	syscallNames[253] = "inotify_init"
	syscallNames[254] = "inotify_add_watch"
	syscallNames[255] = "inotify_rm_watch"
	syscallNames[256] = "migrate_pages"
	syscallNames[257] = "openat"
	syscallNames[258] = "mkdirat"
	syscallNames[259] = "mknodat"
	syscallNames[260] = "fchownat"
	syscallNames[261] = "futimesat"
	syscallNames[262] = "newfstatat"
	syscallNames[263] = "unlinkat"
	syscallNames[264] = "renameat"
	syscallNames[265] = "linkat"
	syscallNames[266] = "symlinkat"
	syscallNames[267] = "readlinkat"
	syscallNames[268] = "fchmodat"
	syscallNames[269] = "faccessat"
	syscallNames[270] = "pselect6"
	syscallNames[271] = "ppoll"
	syscallNames[272] = "unshare"
	syscallNames[273] = "set_robust_list"
	syscallNames[274] = "get_robust_list"
	syscallNames[275] = "splice"
	syscallNames[276] = "tee"
	syscallNames[277] = "sync_file_range"
	syscallNames[278] = "vmsplice"
	syscallNames[279] = "move_pages"
	syscallNames[280] = "utimensat"
	syscallNames[281] = "epoll_pwait"
	syscallNames[282] = "signalfd"
	syscallNames[283] = "timerfd_create"
	syscallNames[284] = "eventfd"
	syscallNames[285] = "fallocate"
	syscallNames[286] = "timerfd_settime"
	syscallNames[287] = "timerfd_gettime"
	syscallNames[288] = "accept4"
	syscallNames[289] = "signalfd4"
	syscallNames[290] = "eventfd2"
	syscallNames[291] = "epoll_create1"
	syscallNames[292] = "dup3"
	syscallNames[293] = "pipe2"
	syscallNames[294] = "inotify_init1"
	syscallNames[295] = "preadv"
	syscallNames[296] = "pwritev"
	syscallNames[297] = "rt_tgsigqueueinfo"
	syscallNames[298] = "perf_event_open"
	syscallNames[299] = "recvmmsg"
	syscallNames[300] = "fanotify_init"
	syscallNames[301] = "fanotify_mark"
	syscallNames[302] = "prlimit64"
	syscallNames[303] = "name_to_handle_at"
	syscallNames[304] = "open_by_handle_at"
	syscallNames[305] = "clock_adjtime"
	syscallNames[306] = "syncfs"
	syscallNames[307] = "sendmmsg"
	syscallNames[308] = "setns"
	syscallNames[309] = "getcpu"
	syscallNames[310] = "process_vm_readv"
	syscallNames[311] = "process_vm_writev"
	syscallNames[312] = "kcmp"
	syscallNames[313] = "finit_module"
	syscallNames[314] = "sched_setattr"
	syscallNames[315] = "sched_getattr"
	syscallNames[316] = "renameat2"
	syscallNames[317] = "seccomp"
	syscallNames[318] = "getrandom"
	syscallNames[319] = "memfd_create"
	syscallNames[320] = "kexec_file_load"
	syscallNames[321] = "bpf"
	syscallNames[322] = "execveat"
	syscallNames[323] = "userfaultfd"
	syscallNames[324] = "membarrier"
	syscallNames[325] = "mlock2"
	syscallNames[326] = "copy_file_range"
	syscallNames[327] = "preadv2"
	syscallNames[328] = "pwritev2"
	syscallNames[329] = "pkey_mprotect"
	syscallNames[330] = "pkey_alloc"
	syscallNames[331] = "pkey_free"
	syscallNames[332] = "statx"
	syscallNames[333] = "io_pgetevents"
	syscallNames[334] = "rseq"
}
