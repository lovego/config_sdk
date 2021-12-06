

var headers = {
  'Accept': 'application/json'
};
var request = http.Request('GET', Uri.parse('http://127.0.0.1:3000/config/pull?project=erp&env=dev&version=1.0&endPointType=server&secret=123&hash='));

request.headers.addAll(headers);

http.StreamedResponse response = await request.send();

if (response.statusCode == 200) {
  print(await response.stream.bytesToString());
}
else {
  print(response.reasonPhrase);
}

