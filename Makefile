GOCMD=go build -v

rft:
	$(GOCMD) -o cbsd-raftd .
	strip cbsd-raftd

clean:
	@rm -f cbsd-raftd

.PHONY: helloworld multigroup ondisk optimistic-write-lock clean
