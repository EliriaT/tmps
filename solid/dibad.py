class XMLHttpService(XMLHttpRequestService):
    pass

class Http:
    def __init__(self, xml_http_service: XMLHttpService):
        self.xml_http_service = xml_http_service
    
    def get(self, url: str, options: dict):
        self.xml_http_service.request(url, 'GET')

    def post(self, url, options: dict):
        self.xml_http_service.request(url, 'POST')


