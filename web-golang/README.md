# Go로 만드는 웹

* json으로 송수신 하는 이유
binary로 주고 받으면 규격을 맞춰야하기 때문에.
그래서 보통 string형태의 데이터 포맷으로 주고 받는 경우가 많다.
원래는 XML로 주고 받았는데, `<tag></tag>`를 열고 닫고 하다보니까 용량이 커져서
그냥 단순한 JavaScript-Object-Notation 형태로 key-value로 하는게 용량도 적고 
간편해서 쓰게 되었다.

하지만, 각 언어에서 이 JSON format을 파싱하는게 꽤나.. 귀찮은 일 ㅠ
