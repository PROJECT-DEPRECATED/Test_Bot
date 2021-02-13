build_path="build/"
version="0.0.1v"

if [ ! -d $build_path ]; then
  echo "$build_path is not exist!"
else
  rm -rf $build_path
  echo "Successful remove path: $build_path"
fi

if [ "$1" == "clean" ]; then
  echo "Clean complete"
else
  rm -rf $build_path
  go build -o TestBot-$version src/TestBot.go

  mkdir $build_path
  mv TestBot-$version $build_path
fi