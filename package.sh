#!/bin/sh
set -e

# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

frp_version=`./bin/frps --version`
echo "build version: $frp_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64 loong64'
extra_all='_ hf'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        for extra in $extra_all; do
            suffix="${os}_${arch}"
            if [ "x${extra}" != x"_" ]; then
                suffix="${os}_${arch}_${extra}"
            fi
            frp_dir_name="ccc_${frp_version}_${suffix}"
            frp_path="./packages/ccc_${frp_version}_${suffix}"

            if [ "x${os}" = x"windows" ]; then
                if [ ! -f "./cccc_${os}_${arch}.exe" ]; then
                    continue
                fi
                if [ ! -f "./ssss_${os}_${arch}.exe" ]; then
                    continue
                fi
                mkdir ${frp_path}
                mv ./cccc_${os}_${arch}.exe ${frp_path}/cccc.exe
                mv ./ssss_${os}_${arch}.exe ${frp_path}/ssss.exe
            else
                if [ ! -f "./cccc_${suffix}" ]; then
                    continue
                fi
                if [ ! -f "./ssss_${suffix}" ]; then
                    continue
                fi
                mkdir ${frp_path}
                mv ./cccc_${suffix} ${frp_path}/cccc
                mv ./ssss_${suffix} ${frp_path}/ssss
            fi  
            cp ../LICENSE ${frp_path}
            cp -f ../conf/cccc.toml ${frp_path}
            cp -f ../conf/ssss.toml ${frp_path}

            # packages
            cd ./packages
            if [ "x${os}" = x"windows" ]; then
                zip -rq ${frp_dir_name}.zip ${frp_dir_name}
            else
                tar -zcf ${frp_dir_name}.tar.gz ${frp_dir_name}
            fi  
            cd ..
            rm -rf ${frp_path}
        done
    done
done

cd -
