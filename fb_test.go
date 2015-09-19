package fb

import (
	"testing"

	"github.com/daaku/go.h"
	"github.com/facebookgo/ensure"
	"golang.org/x/net/context"
)

func TestInit(t *testing.T) {
	i := &Init{
		AppID:   42,
		Version: "2.0.0",
		XFBML:   true,
	}
	s, err := h.Render(context.Background(), i)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, s,
		`<script src="//connect.facebook.net/en_US/all.js" `+
			`async></script><script>window.fbAsyncInit=`+
			`function(){FB.init({"appId":42,"version":"2.0.0",`+
			`"xfbml":true})</script>`)
}
