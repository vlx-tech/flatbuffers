cmake_minimum_required(VERSION 3.0)

fips_begin_lib(flatbuffers)

    include_directories(../include)

    fips_dir(.. GROUP .)
    fips_files(
        include/flatbuffers/code_generators.h
        include/flatbuffers/base.h
        include/flatbuffers/flatbuffers.h
        include/flatbuffers/hash.h
        include/flatbuffers/idl.h
        include/flatbuffers/util.h
        include/flatbuffers/reflection.h
        include/flatbuffers/reflection_generated.h
        include/flatbuffers/stl_emulation.h
        include/flatbuffers/flexbuffers.h
        include/flatbuffers/registry.h
        include/flatbuffers/minireflect.h
        src/code_generators.cpp
        src/idl_parser.cpp
        src/idl_gen_text.cpp
        src/reflection.cpp
        src/util.cpp
        )

fips_end_lib()
