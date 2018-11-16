#!/bin/bash

msg="\
Halo, saya $CIRCLE_USERNAME mau meminta review untuk PR berikut
Github   : $CIRCLE_PULL_REQUEST
<!channel>
"

if [ "$CIRCLE_BRANCH" != "ggzzzz" ] && [ "$CIRCLE_BRANCH" != "staging" ] ; then
    curl -X POST --data-urlencode "payload={\"channel\": \"#testrobby2\", \"username\": \"Robby\", \"text\": \"$msg\",\"link_names\": 1, \"icon_emoji\": \":davidjenggot:\"}" https://hooks.slack.com/services/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
    echo "\nalready slack"
else
    echo "no slack"
fi
