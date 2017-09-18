call gengen.bat mvc -override=true -controller=App -projectPath=um -customPath=/self -root specs
call gengen.bat db -override=true -root specs -output=app/models
call gengen.bat test_base -override=true -projectPath=um -root specs

FOR %%i IN (app\controllers\*.go) DO goimports -w %%i
FOR %%i IN (app\models\*.go) DO goimports -w %%i
FOR %%i IN (tests\*.go) DO goimports -w %%i