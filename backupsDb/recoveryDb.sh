#!/bin/bash

. /PathTo/settings.sh

info="Recovery is connected"
pg_restore  -U $db_user  -d $db_name  -F d $db_dump_catalog_path