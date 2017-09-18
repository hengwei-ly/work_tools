
if "%gengen_path%" EQU "" (
  if exist D:\developing\go\meijing\tpt_vendor\src\github.com\three-plus-three\gengen (
    set gengen_path=D:\developing\go\meijing\tpt_vendor\src\github.com\three-plus-three\gengen
  )
  if exist D:\go\tpt_vendor\src\github.com\three-plus-three\gengen (
    set gengen_path=D:\go\tpt_vendor\src\github.com\three-plus-three\gengen
  )
)

@pushd %gengen_path%
@%gengen_path%\gengen.exe embeded template_gen.go
@goimports  -w template_gen.go
@go build
@%gengen_path%\gengen.exe embeded template_gen.go
@goimports  -w template_gen.go
@go build
@popd
