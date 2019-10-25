# Makefile

DOTFILES = $(wildcard *.dot)
PNGFILES = $(DOTFILES:.dot=.png)
TARGET = $(PNGFILES)

all: $(TARGET)

%.ps: %.dot
	dot -Tps $^ -o $@

%.svg: %.dot
	dot -Tsvg $^ -o $@

%.png: %.dot
	dot -Tpng $^ -o $@

clean:
	rm -f $(TARGET)

