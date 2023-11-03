.PHONY: build
build:
	tinygo build --target wioterminal -o build/blinky.uf2
	
.PHONY: flash
flash:
	tinygo flash --target wioterminal

.PHONY: monitor
monitor:
	tinygo monitor --target wioterminal

.PHONY: flash-monitor
flash-monitor:
	tinygo flash --target wioterminal --monitor

.PHONY: minicom
minicom:
	minicom -D /dev/tty.usbmodem101
