all: builtin pdf

builtin:
	pandoc --self-contained --metadata pagetitle="Manual" --css builtin.css -o ../cmd/gui/res/help.html manual.md

pdf:
	grep -v HIDEBUILTIN manual.md | pandoc -s --pdf-engine=xelatex -o manual.pdf

clean:
	rm -f manual.pdf ../cmd/gui/res/help.html