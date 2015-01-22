#!/bin/bash
for ((i=1;i<10000;i++)); do
    curl localhost:8080/users -i -d '{"name":"chuan 
    ldjfldsalkfjldsajklfajlfjsafasld
    jldsafjklsajlfjsa;lkjf
    jlfdsakjf;saf
    jflkdsjfsa;
    jlsfdksajoiewo
    jfsaldajkfajfa
    jflkdsajf
    jldsakf;jajdsieowjf;lsajdflka
    <F5>ajldjlajljjafdjdlsajiej d sajljfjafjldsajie"}' >/dev/null 2>&1
done
