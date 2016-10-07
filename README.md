
The R and Golang scripts demo'd at the NH Data Science Meetup

Here's where I pulled the bits out for the R script: [link](https://rstudio-pubs-static.s3.amazonaws.com/31867_8236987cf0a8444e962ccd2aec46d9c3.html)

Here's the web page where you play around with some of the basic Golang code: [link](https://play.golang.org/)

I was hoping to have time to try out this stemmer as well
[Golang Snowball stemmer](https://github.com/kljensen/snowball)

And some resources for text files:

* [British Law (used in demo)](https://www.gutenberg.org/wiki/British_Law_(Bookshelf))
* [Project Gutenberg](https://www.gutenberg.org/)
* [Newsgroups](http://qwone.com/~jason/20Newsgroups/)

My lazy man's "make small data into bigger data script"

```bash
#!/bin/bash

DIR=british-law

for N in $(seq 0 1000); do
TIME=$(date +%s)
    for FILE in $(ls $DIR/*.txt); do
        cp $FILE{,.$TIME}
    done
done
```

