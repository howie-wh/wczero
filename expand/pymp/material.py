# -*- coding: utf-8 -*-
# filename: material.py
import urllib.request
import json
import poster3.encode
from poster3.streaminghttp import register_openers
from basic import Basic

class Material(object):
    def __init__(self):
        register_openers()
    #上传
    def upload(self, accessToken, filePath, mediaType):
        openFile = open(filePath, "rb")
        fileName = "hello"
        param = {'media': openFile, 'filename': fileName}
        #param = {'media': openFile}
        postData, postHeaders = poster.encode.multipart_encode(param)
        if isinstance(postData, str):
                    postData = postData.encode('utf-8')

        postUrl = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s" % (accessToken, mediaType)
        urlResp = urllib.request.urlopen(url=postUrl, data=postData)
        print(urlResp.read())
    #下载
    def get(self, accessToken, mediaId):
        postUrl = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s" % accessToken
        postData = "{ \"media_id\": \"%s\" }" % mediaId
        if isinstance(postData, str):
                    postData = postData.encode('utf-8')
        urlResp = urllib.request.urlopen(url=postUrl, data=postData)
        headers = urlResp.info().__dict__['headers']
        if ('Content-Type: application/json\r\n' in headers) or ('Content-Type: text/plain\r\n' in headers):
            jsonDict = json.loads(urlResp.read())
            print(jsonDict)
        else:
            buffer = urlResp.read()  # 素材的二进制
            mediaFile = file("test_media.jpg", "wb")
            mediaFile.write(buffer)
            print("get successful")
    #删除
    def delete(self, accessToken, mediaId):
        postUrl = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s" % accessToken
        postData = "{ \"media_id\": \"%s\" }" % mediaId
        if isinstance(postData, str):
                    postData = postData.encode('utf-8')
        urlResp = urllib.request.urlopen(url=postUrl, data=postData)
        print(urlResp.read())

    #获取素材列表
    def batch_get(self, accessToken, mediaType, offset=0, count=20):
        postUrl = ("https://api.weixin.qq.com/cgi-bin/material"
               "/batchget_material?access_token=%s" % accessToken)
        postData = "{ \"type\": \"%s\", \"offset\": %d, \"count\": %d }" % (mediaType, offset, count)
        if isinstance(postData, str):
            postData = postData.encode('utf-8')
        urlResp = urllib.request.urlopen(url=postUrl, data=postData)
        print(urlResp.read())

if __name__ == '__main__':
    myMaterial = Material()
    accessToken = Basic().get_access_token()
    mediaType = "news"
    myMaterial.batch_get(accessToken, mediaType)
