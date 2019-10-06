<!DOCTYPE html>
<html>
<body>
 
<script>
	var rozsah = 345
	Array.prototype.timeoutSort = function (f) {
	  this.forEach((n) => setTimeout( () => f(n) , Math.PI * (rozsah - n)));
	}
	var pole = [];
	for(var i = 0; i < Math.random()*1234; i++) {
	  pole[Math.floor(Math.random()*2345)]=Math.floor(Math.random()*rozsah);
	}
	pole.timeoutSort(function(n) { document.write(n + ' '); });
</script>
 
</body>
</html>
