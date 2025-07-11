package hello

import (
    "rsc.io/quote"
    "rsc.io/sampler"
)

func Hello() string {
    sampler.Hello()
    return quote.Hello()
}


