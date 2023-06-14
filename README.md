# 3.MicroserviceEngineering

1. Execute main.go
2. Call POST localhost:8080/cars to create a new car. Example Payload:
<code>
{
   "vin": "1",
   "brand": "Mercedes Benz",
   "model": "S Klasse"
}
</code>
3. Call GET localhost:8080/cars/<vin> to retrieve created car. Example vin <code>1</code>
