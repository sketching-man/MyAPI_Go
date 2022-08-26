# Web Crawler란?

SEED URL에서 시작하여 관련된 URL을 찾아 내고, 그 URL들에서 또 다른 하이퍼 링크를 계속 하여 찾아내는 과정을 반복하며 하이퍼 링크들을 다운로드하는 프로그램.

# Web Crawler 구조

![](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FEMfB5%2FbtrblP1weEB%2FEFYoRN7GzslC66NmmuEqa1%2Fimg.jpg)

* Frontier: 2개의 자료구조(탐색했던 URL과, 탐색해야 할 URL)를 바탕으로 탐색할 URL에 하나 이상의 Seed URL을 Fetcher 에게 전달
* Fetcher: URL의 html 내용을 가져와서 Parser에게 전달
* Parser: html에서 \<a> 태그 혹은 \<url> 태그의 다른 하이퍼링크를 탐색
* Content Seen: 방문한 페이지의 본문 내용이 이미 전에 본 내용인지 아닌지 판단
(Doc FP`s = Document Finger Prints) = Document의 내용을 구분할 수 있는 고유 Hash 값
* URL Filter: 해당 html 페이지에서 \<a> 태그 혹은 \<html> 태그 등의 URL을 분류
* Dup URL Elim: 방문한 페이지를 다시 방문하는 LOOP가 발생하지 않도록, 중복된 URL을 제거하여 Frontier에게 전달

## 문서 중복 판단 방법

https://github.com/aragorn/home/wiki/Near-Duplicate-Documents-Detection

# robots.txt

Crawler가 페이지에 요청을 해도 되거나, 요청을 해서는 안되는 경로를 적어둔 File.  
계속적인 접근으로 인한 과부하 혹은 적당하지 않은 정보를 가져가는 것을 막기 위한 규약들이 기록.

ex) https://www.google.com/robots.txt

이러한 robots.txt에 정의된 규약들을 지키지 않고 데이터를 추출하거나, 추출한 데이터를 통해 이익을 취하는 Crawler는 불법 Crawler라고 판단.

# 참고 자료

https://velog.io/@mowinckel/%EC%9B%B9-%ED%81%AC%EB%A1%A4%EB%A7%81-I
https://soyoung-new-challenge.tistory.com/93
https://well-made-codestory.tistory.com/21