set -o errexit
set -o nounset

export PARSER_CONTAINER=go-parser

if [ ! -d "./parser_results" ]
then
  echo "Directory not exists. Creating ./parser_results"
  mkdir ./parser_results
  echo "Directory created!"
fi

echo "Copy file to: ./parser_results"
winpty docker cp $PARSER_CONTAINER:/parser/out.json d:\collocutor\project\parser\parser_results
cat "./parser_results/out.json"
echo "file copied."