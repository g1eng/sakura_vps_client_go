#!/bin/bash -ex
# This script make temporary fixation for the original spec,
# which lacks pagenation functionality for API key-related operations.

SRC=${1:?no source file specified}
DST=${2:?no destination file specified}


{
  l=1
  while read -r line ; do
  
  case $l in
  2532|2815|3001) echo '
            {
              "$ref": "#/components/parameters/PaginationPage"
            },
            {
              "$ref": "#/components/parameters/PaginationPerPage"
            },'
   ;;
  esac
  echo "$line"
  l=$((l+1))
  
  done < $SRC
  echo "}"
} > $DST
