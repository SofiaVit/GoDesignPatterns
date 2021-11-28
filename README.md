# Go Design Patterns
<details><summary> Abstract Factory </summary>
The abstract factory pattern provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes. </br>
<h4> Abstract Factory - Restaurant example </h4>
The Abstract Factory restaurant can create only two types of meals, Italian and American. Each meal consists of a dish and drink. </br>
When a client orders a meal from a certain cuisine type, he should get a drink and a dish, both from the same cuisine type he chooses.
</br></br>
<a href="https://github.com/SofiaVit/GoDesignPatterns/tree/master/abstractfactory">
Link to example code</a></br>
<a href="https://github.com/SofiaVit/GoDesignPatterns/blob/master/uml/GoAbstractFactoryUML.png">
Link to example UML</a></br>
</details>

<details><summary> Factory </summary>
The factory pattern uses factory methods to deal with the problem of creating objects without having to specify the exact class of the object that will be created.</br>
<h4> Factory - Delivery Service example </h4>
The delivery service needs to choose a deleviry vehicle and schedule it based on the distance from the client.
If the distance is less than 3.5 k"m it will use bicycle, between 3.5 k"m and 25 k"m it will use car and wont deliver to distance more than 25 k"m.
</br></br>
<a href="https://github.com/SofiaVit/GoDesignPatterns/tree/master/factory">
Link to example code</a></br>
<a href="">
Link to example UML</a></br>
</details>
