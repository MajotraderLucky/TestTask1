#!/bin/bash

. /PathTo/settings.sh
. /PathTo/recoveryDb.sh
echo "$createDump"
echo "$recoveryDB"
echo "$info"

echo "pg_dump -U $db_user -F d -f $db_dump_catalog_path $db_name -j $db_streams $db_schema"
pg_dump -U $db_user -F d -f $db_dump_catalog_path $db_name -j $db_streams $db_schema