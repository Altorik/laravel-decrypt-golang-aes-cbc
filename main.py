from Crypto.Cipher import AES
import base64
import demjson
import datetime
from phpserialize import serialize, unserialize


class decrypter:
    def __init__(self, key):
        self.key = base64.b64decode(key)

    def decrypt(self, text):
        decoded_text = demjson.decode(base64.b64decode(text))
        iv = base64.b64decode(decoded_text['iv'])
        crypt_object = AES.new(key=self.key, mode=AES.MODE_CBC, IV=iv)
        decoded = base64.b64decode(decoded_text['value'])
        decrypted = crypt_object.decrypt(decoded)
        return str(unserialize(decrypted)).strip()


begin_time = datetime.datetime.now()
# Key from .env file
obj = decrypter(key_from_env)
# text to decode
alex = obj.decrypt(text_to_decode)
print(alex)
print(datetime.datetime.now() - begin_time)
