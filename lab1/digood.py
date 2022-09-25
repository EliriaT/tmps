class Connection:
    def request(self, url: str, options: dict):
        raise NotImplementedError


class Http:
    def __init__(self, http_connection: Connection):
        self.http_connection = http_connection
    
    def get(self, url: str, options: dict):
        self.http_connection.request(url, 'GET')

    def post(self, url, options: dict):
        self.http_connection.request(url, 'POST')


class RestHttpService(Connection):
    

    def request(self, url: str, options:dict):
      pass

class NodeHttpService(Connection):
    def request(self, url: str, options:dict):
        pass

class HttpsService(Connection):
    def request(self, url: str, options:dict):
        pass

