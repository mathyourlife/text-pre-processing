# install.packages("devtools")
# install.packages("plyr")
# install.packages("ggplot2")
# install.packages("wordcloud")
# install.packages("tm")
# install.packages("SnowballC")

library(tm)
library(SnowballC)
setwd("~/gocode/src/github.com/mathyourlife/text-processing")
path = "data/www.gutenberg.org/british-law-large"

blaw_dir = DirSource(path, encoding = "UTF-8")
blaw_corpus_raw = Corpus(blaw_dir)
blaw_corpus_raw[[1]]$content

blaw_corpus = blaw_corpus_raw
blaw_corpus <- tm_map(blaw_corpus, removePunctuation)   # *Removing punctuation:*    
blaw_corpus <- tm_map(blaw_corpus, removeNumbers)      # *Removing numbers:*    
blaw_corpus <- tm_map(blaw_corpus, tolower)   # *Converting to lowercase:*    
blaw_corpus <- tm_map(blaw_corpus, removeWords, stopwords("english"))   # *Removing "stopwords" 

blaw_corpus[[1]]

blaw_corpus <- tm_map(blaw_corpus, stemDocument)   # *Removing common word endings* (e.g., "ing", "es")   
blaw_corpus <- tm_map(blaw_corpus, stripWhitespace)   # *Stripping whitespace   
blaw_corpus <- tm_map(blaw_corpus, PlainTextDocument) 

blaw_corpus[[1]]$content

## *This is the end of the preprocessing stage.*   


### Stage the Data      
dtm <- DocumentTermMatrix(blaw_corpus)   
# tdm <- TermDocumentMatrix(blaw_corpus)

### Word Clouds!   
# First load the package that makes word clouds in R.    
library(wordcloud)
dtms <- removeSparseTerms(dtm, 0.15) # Prepare the data (max 15% empty space)   
freq <- colSums(as.matrix(dtm)) # Find word frequencies   
dark2 <- brewer.pal(6, "Dark2")   

png("british-law-cloud.png", width=1280,height=800)
wordcloud(names(freq), freq, max.words=100, rot.per=0.2, colors=dark2)   
dev.off()
