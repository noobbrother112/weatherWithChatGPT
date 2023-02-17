## 현재 날씨기반 옷차림 추천 사이트



접속자의 ip를 기반으로 유저 접속위치 추정, 근처의 날씨를 보여주며 <br>(skt의 경우 모든 아이피의 위치가 서울특별시 중구 을지로65로 표기되는 이슈가 있음 - 추후 위치기반 서비스로 변경 예정)<br>
 
(서울기준) 날씨를 chatGPT-3모델에 보내 추천하는 의상을 받아 보여줌




### 준비물    
>chatGPT Api 키 <br>
```export API_KEY={YOUR_KEY}```<br><br>
공공데이터포털 whois api 키   <br>
```export WHOIS_API_KEY={YOUR_KEY}```<br><br>
공공데이터포털 기상청 중기예보 api 키<br>
```export WEATHER_API_KEY={YOUR_KEY}```<br> 

<br>// chatGPT의 기능이 궁금하여 시작한 토이프로젝트이며 
작업중 chatGPT를 어느정도 활용하여 코드를 짠 후 수정하는 방향으로 작업 했고, 보조적인 역할에서의 성능에 감탄을 금치 못하였음