#!/bin/bash

msg="\
Halo, @viriyahendarta
Username : $CIRCLE_USERNAME
Github   : $CIRCLE_PULL_REQUEST

<!channel>
"



curl -X POST --data-urlencode "payload={\"channel\": \"#testrobby2\", \"username\": \"Robby\", \"text\": \"$msg\",\"link_names\": 1, \"icon_emoji\": \":davidjenggot:\"}" https://hooks.slack.com/services/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