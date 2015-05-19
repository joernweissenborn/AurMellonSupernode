package boot

func Boot(){
	n := BootAurArath()
	in := BootAurMellon(n)
	t := BootTransactor()
	BootProcessors(n,in,t)
	n.Run()
}
