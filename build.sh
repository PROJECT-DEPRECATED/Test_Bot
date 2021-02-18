build_path="build/"
version="0.0.2v"

if [ ! -d $build_path ]; then
  echo "$build_path is not exist!"
else
  rm -rf $build_path
  echo "Successful remove path: $build_path"
fi

if [ "$1" == "clean" ]; then
  echo "Clean complete"

elif [ "$1" == "run" ]; then
  # You can run this bot with token.txt
  # shellcheck disable=SC2034
  token=$(cat token.txt)
  # shellcheck disable=SC2086
  go run src/TestBot.go -token $token


else
  rm -rf $build_path
  go build -o TestBot-$version src/TestBot.go

  mkdir $build_path
  mv TestBot-$version $build_path
fi