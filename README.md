# BK-tree

There is some sources with BK-trees:

* [signal-to-noise.xyz (lang: en)](http://signal-to-noise.xyz/post/bk-tree/)
* [blog.mishkovskyi.net (lang: en)](http://blog.mishkovskyi.net/posts/2015/Jul/31/implementing-bk-tree-in-clojure)
* [geeksforgeeks (lang: en)](http://www.geeksforgeeks.org/bk-tree-introduction-implementation/)
* [habrahabr.ru (lang: ru)](https://habrahabr.ru/post/114997/)
* [t2a.co (lang: en)](https://t2a.co/blog/index.php/spell-checking-or-search-engine-suggestions-using-bk-trees/)
* [kafsemo.org (lang: en)](http://www.kafsemo.org/2010/08/03_bk-tree-performance-notes.html)
* [kallistec.wordpress.com (lang: en)](https://kallistec.wordpress.com/tag/k-nearest-neighbors/)


In addition to my implementation, there are several more:

* [agatan (any objects, any metrics, only int distance, use interface)](https://github.com/agatan/bktree/blob/master/bktree.go)
* [mukulrawat1986 (only string, only levenshtein, only int distance)](https://github.com/mukulrawat1986/bktree-go/blob/master/bktree/bktree.go)
* [gansidui (only string, only levenshtein, only int distance)](https://github.com/gansidui/bktree/blob/master/bktree.go)
* [maccman (any objects, any metrics, only int distance, use calback function for metrics)](https://github.com/maccman/bktree/blob/master/bktree.go)
* [hjr265](https://github.com/hjr265/go-bktree/blob/master/bktree.go)
* [umahmood](https://github.com/umahmood/bktree/blob/master/bktree.go)
* [vy](https://github.com/vy/bk-tree)

My implementation was made [agatan](https://github.com/agatan/bktree/blob/master/bktree.go)-style code. I modified these sources and made more speedly used map. In my implementation there is abstract type for distance (float32, int and your other types) in front of [agatan](https://github.com/agatan/bktree/blob/master/bktree.go) implementation.

These trees can be used for fuzzy information retrieval, as well as for other data analysis tasks. For example, these trees are closely related to the [kNN algorithm](https://kallistec.wordpress.com/tag/k-nearest-neighbors/).

It is also worth noting that these trees can be generalized using a probabilistic model and used as a spell checker. To do this, it is necessary to select words in the radius R from the given one and take into account the proximity of other words to the ones found. Plus, this will take into account the frequency of words in the investigated case of documents (if it is available).

To start the project, you need to set dependencies with [glide](https://glide.readthedocs.io/en/latest/). Steps:

* set environment `$GOPATH`

* download packages `go get bitbucket.org/egorsteam/bk-tree`

To install dependencies you need:

* execute `curl https://glide.sh/get | sh`

* execute `go get -u github.com/Masterminds/glide`

* execute in work directory `glide install`

Regarding BK trees, I have some misunderstanding. Obviously, having two vertices belonging to the same branch, we can easily say that:

![f1](https://latex.codecogs.com/gif.latex?%5Crho%28A%2C%20B%29%20%5Cleq%20%5Crho%28A%2C%20C%29%20&plus;%20%5Crho%28C%2C%20B%29)

On the other hand, if A, B, or C pairwise belong to different branches of different subtrees, then will the triangle inequality hold or not?

# BK-деревья

Об этих деревьях написано в нескольких источниках:

* [signal-to-noise.xyz (язык: англ.)](http://signal-to-noise.xyz/post/bk-tree/)
* [blog.mishkovskyi.net (язык: англ.)](http://blog.mishkovskyi.net/posts/2015/Jul/31/implementing-bk-tree-in-clojure)
* [geeksforgeeks (язык: англ.)](http://www.geeksforgeeks.org/bk-tree-introduction-implementation/)
* [habrahabr.ru (язык: ру.)](https://habrahabr.ru/post/114997/)
* [t2a.co (язык: англ.)](https://t2a.co/blog/index.php/spell-checking-or-search-engine-suggestions-using-bk-trees/)
* [kafsemo.org (язык: англ.)](http://www.kafsemo.org/2010/08/03_bk-tree-performance-notes.html)
* [kallistec.wordpress.com (язык: англ.)](https://kallistec.wordpress.com/tag/k-nearest-neighbors/)


Помимо моей реализации, есть ещё несколько:

* [agatan (any objects, any metrics, only int distance, use interface)](https://github.com/agatan/bktree/blob/master/bktree.go)
* [mukulrawat1986 (only string, only levenshtein, only int distance)](https://github.com/mukulrawat1986/bktree-go/blob/master/bktree/bktree.go)
* [gansidui (only string, only levenshtein, only int distance)](https://github.com/gansidui/bktree/blob/master/bktree.go)
* [maccman (any objects, any metrics, only int distance, use calback function for metrics)](https://github.com/maccman/bktree/blob/master/bktree.go)
* [hjr265](https://github.com/hjr265/go-bktree/blob/master/bktree.go)
* [umahmood](https://github.com/umahmood/bktree/blob/master/bktree.go)
* [vy](https://github.com/vy/bk-tree)

Моя реализация сделана в стиле [agatan](https://github.com/agatan/bktree/blob/master/bktree.go). Я модифицировал его код и сделал его чуть быстрее при помощи map. В отличие от реализации [agatan](https://github.com/agatan/bktree/blob/master/bktree.go), в моей есть абстрактный тип расстояния (float32, int и любой другой Ваш тип).

Эти деревья можно использовать для нечёткого информационного поиска, а также для иных задач анализа данных. Так, например, эти деревья тесно связаны с [алгоритмом kNN](https://kallistec.wordpress.com/tag/k-nearest-neighbors/).

Также стоит заметить, что данные деревья можно обобщить с использованием вероятностной модели и использовать как spell checker. Для этого необходимо выбирать слова в радиусе R от заданного и учитывать близость других слов к найденным. Плюсом к этому будет учёт частоты слов в исследуемом корпусе документов (если он есть в наличии).

В отношении BK-деревьев у меня есть некоторое непонимание. Очевидно, что имея две вершины, принадлежащие одной и той же ветке, мы легко можем сказать, что:

![f1](https://latex.codecogs.com/gif.latex?%5Crho%28A%2C%20B%29%20%5Cleq%20%5Crho%28A%2C%20C%29%20&plus;%20%5Crho%28C%2C%20B%29)

С другой стороны, если A, B или C попарно принадлежат разным веткам различных поддеревьев, то в таком случае будет выполняться неравенство треугольника или нет?

Чтобы запустить пакеты, необходимо установить [glide](https://glide.readthedocs.io/en/latest/). Шаги:

* Установить переменную окружения `$GOPATH`

* Скачайте пакет с исходниками `go get bitbucket.org/egorsteam/bk-tree`

Чтобы разрешить зависимости, Вам нужно:

* Выполните `curl https://glide.sh/get | sh`

* Выполните `go get -u github.com/Masterminds/glide`

* В рабочей директории пакета выполните `glide install`