# Exercise 4.10
"Modify [issues](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch4/issues/main.go) to report the results in age categories, say less than a month old, less than a year old, and more than a year old."

# Results
```
$ go run main.go repo:golang/go cmd asm
657 issues:
Less than a month old
#26110 cherrymui cmd/asm: reject use of R18 on ARM64?
#26469       rsc cmd/compile: short tests are not short enough
#26141     mvdan runtime: SIGSEGV on a freebsd-amd64 builder
#26517     gooid cmd/cgo: crash when building github.com/timob/jnigi on

Less than a year old
#25724   robpike cmd/asm: how to test?
#25274  chai2010 cmd/asm: allow methods implemented in assembly
#21860 Quasilyte cmd/internal/obj/x86: improve asm error messages
#25770   artooro runtime: SIGSEGV when building on freebsd ARM
#23386 Quasilyte cmd/internal/obj/x86: FSAVE assembled as FNSAVE
#23274 alllexx88 cmd/dist: detect FPUless ARMv7 platforms?
#22333 jimrobins runtime: signal 28 received on thread with no signal st
#25766 FiloSotti x/build/cmd/gitmirror: unexpected signal
#21735  martisch cmd/compile: avoid slow versions of LEA instructions on

More than a year old
#17544 aclements cmd/vet: move asm checks into cmd/asm
#14069       rsc cmd/asm: incorrect instruction encodings
#21004  crawshaw cmd/asm, runtime: textflag for CABI
#14288 randall77 cmd/asm: accepts MOVB (BX)(R8*4), AH incorrectly
#18072   robpike cmd/asm/internal/asm: add encoding tests for all archit
#12655   dvyukov cmd/asm: nil deref in LOOP (PC)
#7300  gopherbot cmd/asm: add neon, vector instructions for arm
#4978  MichaelTJ cmd/compile: allow methods implemented in assembly
#20173  dlespiau cmd/asm: change canonical spelling of CMOVLEQ to CMOVL.
#17235   laboger cmd/compile, cmd/internal/obj/ppc64: generate the likel
#14028 Oppositio cmd/compile: stack overflow accessing large return valu
#15512   brtzsnr cmd/compile: evaluate the need for shortcircuit
#10870     eloff cmd/asm: Allow data references to labels in assembly (s
#12658   dvyukov cmd/asm: unactionable "invalid local variable type 0"
#13554 pindamonh cmd/compile: runs out of memory allocating 12k + items
#15770   laboger cmd/link: panic: runtime error: invalid memory address
#20982     ghost cmd/link: support msvc object files
```