CXX_X86 := i686-w64-mingw32-gcc
CXX_X64 := x86_64-w64-mingw32-gcc
GO_BUILD_FLAGS := -a -v --gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -buildmode=c-archive 

default:
	mkdir compiled
# Build x86
	CC=$(CXX_X86) CGO_ENABLED=1 GOOS=windows GOARCH=386 go build $(GO_BUILD_FLAGS) -o HackBrowserData.x86.a .
	$(CXX_X86) dllmain.def HackBrowserData.x86.a -shared -lwinmm -lws2_32 -o compiled/HackBrowserData.x86.dll
# Build x64
	CC=$(CXX_X64) CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o HackBrowserData.x64.a .
	$(CXX_X64) dllmain.def HackBrowserData.x64.a -shared -lwinmm -lws2_32 -o compiled/HackBrowserData.x64.dll

clean:
	@rm -rf *.a *.dll compiled HackBrowserData.*.h