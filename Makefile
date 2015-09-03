CMD=go build -v -o 301
GOOS=export GOOS
GOARCH=export GOARCH
LINUX=$(GOOS)=linux
OSX=$(GOOS)=darwin

windows64:
	$(GOOS)=windows && $(GOARCH)=amd64 && go build -v -o 301.exe \
	&& zip windows64.zip 301.exe && rm 301.exe

linux_arm:
	$(LINUX) && $(GOARCH)=arm && $(CMD) \
		&& tar czf linux_arm.tar.gz 301 && rm 301

linux64:
	$(LINUX) && $(GOARCH)=amd64 && $(CMD) \
		&& tar czf linux64.tar.gz 301 && rm 301

linux386:
	$(LINUX) && $(GOARCH)=386 && $(CMD) \
		&& tar czf linux386.tar.gz 301 && rm 301

osx64:
	$(OSX) && $(GOARCH)=amd64 && $(CMD) \
		&& tar czf osx64.tar.gz 301 && rm 301

osx386:
	$(OSX) && $(GOARCH)=386 && $(CMD) \
		&& tar czf osx386.tar.gz 301 && rm 301

all: windows64 linux64 linux386 linux_arm osx64 osx386

