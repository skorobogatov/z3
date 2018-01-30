Z3_REF ?= master

all: Z3Prover/build/libz3.a

clean:
	rm -rf Z3Prover

Z3Prover/build/libz3.a: Z3Prover
	cd Z3Prover && python scripts/mk_make.py --staticlib --noomp
	cd Z3Prover/build && ${MAKE} api_dll

Z3Prover:
	git clone https://github.com/Z3Prover/z3.git Z3Prover
	cd Z3Prover && git reset --hard && git clean -fdx
	cd Z3Prover && git checkout ${Z3_REF}
