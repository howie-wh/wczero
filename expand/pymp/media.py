# -*- coding: utf-8 -*-
# filename: media.py
from basic import Basic
import urllib.request
import poster3.encode
from poster3.streaminghttp import register_openers


class Media(object):
    def __init__(self):
        register_openers()

    # 上传图片
    def upload(self, accessToken, filePath, mediaType):
        openFile = open(filePath, "rb")
        param = {'media': openFile}
        postData, postHeaders = poster3.encode.multipart_encode(param)

        postUrl = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s" % (
            accessToken, mediaType)

        data =urllib.request.urlopen(postUrl, postData, postHeaders).read()
        print(data)


if __name__ == '__main__':
    myMedia = Media()
    # accessToken = Basic().get_access_token()
    accessToken = "56_NM67bYcQ_Bnv4b2U0B-fiDtzlNbiakApR-DkHQoo_tKCxR2ac3tTMdnptf8WW9DsLlVQmfXfuhDqQAlnFpxVJv0KW5vj_" \
                  "SfURUQE3fcUe9ztjK5I1Z7TsrB32lpBJXNMTcQ8GOE9Ck7nR91EOYPbADAGQH"
    filePath = "D:/code/mpGuide/media/test.jpg"  # 请按实际填写
    mediaType = "image"
    myMedia.upload(accessToken, filePath, mediaType)
