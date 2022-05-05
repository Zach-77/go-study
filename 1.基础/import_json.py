#!/usr/bin/env python3
# -*- coding = utf-8 -*-
import json

test =  """{
        "PortalSerialNo": "OPENUBG6000047-20220406-Gdv0KTHYo0OJ",
        "PayChannel": "wechat_ecommerce",
        "message": "ClientIP=127.0.0.1&LocalIP=10.223.7.18&PayChannel=wechat_ecommerce&PortalSerialNo=OPENUBG6000047-20220406-Gdv0KTHYo0OJ&ReqParams=post_params%3D%7B%22UserId%22%3A%22170000040649170%22%2C%22MidasAppId%22%3A%22lotus_001%22%2C%22ProductId%22%3A%22LT012204061639570057530466X%22%2C%22Quantity%22%3A1%2C%22Amt%22%3A20000%2C%22RealPayAmt%22%3A20000%2C%22CurrencyType%22%3A%22CNY%22%2C%22OutTradeNo%22%3A%22LT012204061639570057530466X%22%2C%22Channel%22%3A%22wechat_ecommerce%22%2C%22PayChannelSubId%22%3A2%2C%22TransactionId%22%3A%22E-220406160000404100%22%2C%22Version%22%3A%22v5%22%2C%22ProvideType%22%3A12%2C%22Metadata%22%3A%22%2525257B%25252522oriMchId%25252522%2525253A%252525221623199763%25252522%2525252C%25252522paymentConfigNo%25252522%2525253A%25252522563061882005528577%25252522%2525252C%25252522routeTag%25252522%2525253A%25252522YM%25252522%2525252C%25252522tradeNo%25252522%2525253A%25252522LT012022040616395741188010466%25252522%2525257D%22%2C%22ChannelOrderId%22%3A%224200001322202204062281108149%22%2C%22MidasChannelOrderId%22%3A%2220220406402115%22%2C%22SubOrderList%22%3A%5B%5D%2C%22ChannelExternalUserInfoList%22%3A%5B%5D%2C%22Ts%22%3A1649234806%2C%22Action%22%3A%22NotifyPayment%22%7D%26uni_providetype%3D12%26secret_id%3D1111%26host%3D%26proxy_method%3D&Response=&ResultCode=-1504&ResultInfo=json_parse fail, uni_providetype[12]&Timestamp=2022/04/06 16:46:46 819088&Url=http%3A%2F%2F10.223.7.33%3A80%2Fapi%2Fmtob%2Fpayment_gw%2FpaymentNotify&",
        "ClientIP": "127.0.0.1",
        "Timestamp": "2022/04/06 16:46:46 819088",
        "Url": "http://10.223.7.33:80/api/mtob/payment_gw/paymentNotify",
        "Response": "",
        "@message": "[2022-04-06 16:46:46 819176] INFO ClientIP=127.0.0.1&LocalIP=10.223.7.18&PayChannel=wechat_ecommerce&PortalSerialNo=OPENUBG6000047-20220406-Gdv0KTHYo0OJ&ReqParams=post_params%3D%7B%22UserId%22%3A%22170000040649170%22%2C%22MidasAppId%22%3A%22lotus_001%22%2C%22ProductId%22%3A%22LT012204061639570057530466X%22%2C%22Quantity%22%3A1%2C%22Amt%22%3A20000%2C%22RealPayAmt%22%3A20000%2C%22CurrencyType%22%3A%22CNY%22%2C%22OutTradeNo%22%3A%22LT012204061639570057530466X%22%2C%22Channel%22%3A%22wechat_ecommerce%22%2C%22PayChannelSubId%22%3A2%2C%22TransactionId%22%3A%22E-220406160000404100%22%2C%22Version%22%3A%22v5%22%2C%22ProvideType%22%3A12%2C%22Metadata%22%3A%22%2525257B%25252522oriMchId%25252522%2525253A%252525221623199763%25252522%2525252C%25252522paymentConfigNo%25252522%2525253A%25252522563061882005528577%25252522%2525252C%25252522routeTag%25252522%2525253A%25252522YM%25252522%2525252C%25252522tradeNo%25252522%2525253A%25252522LT012022040616395741188010466%25252522%2525257D%22%2C%22ChannelOrderId%22%3A%224200001322202204062281108149%22%2C%22MidasChannelOrderId%22%3A%2220220406402115%22%2C%22SubOrderList%22%3A%5B%5D%2C%22ChannelExternalUserInfoList%22%3A%5B%5D%2C%22Ts%22%3A1649234806%2C%22Action%22%3A%22NotifyPayment%22%7D%26uni_providetype%3D12%26secret_id%3D1111%26host%3D%26proxy_method%3D&Response=&ResultCode=-1504&ResultInfo=json_parse fail, uni_providetype[12]&Timestamp=2022/04/06 16:46:46 819088&Url=http%3A%2F%2F10.223.7.33%3A80%2Fapi%2Fmtob%2Fpayment_gw%2FpaymentNotify&",
        "LocalIP": "10.223.7.18",
        "ResultInfo": "json_parse fail, uni_providetype[12]",
        "url_uri": null,
        "ResultCode": "-1504",
        "url_http": "http://10.223.7.33:80/api/mtob/payment_gw/paymentNotify",
        "ReqParams_host": "",
        "ReqParams_secret_id": "1111",
        "ReqParams_post_params": "{\\"UserId\\":\\"170000040649170\\",\\"MidasAppId\\":\\"lotus_001\\",\\"ProductId\\":\\"LT012204061639570057530466X\\",\\"Quantity\\":1,\\"Amt\\":20000,\\"RealPayAmt\\":20000,\\"CurrencyType\\":\\"CNY\\",\\"OutTradeNo\\":\\"LT012204061639570057530466X\\",\\"Channel\\":\\"wechat_ecommerce\\",\\"PayChannelSubId\\":2,\\"TransactionId\\":\\"E-220406160000404100\\",\\"Version\\":\\"v5\\",\\"ProvideType\\":12,\\"Metadata\\":\\"%25257B%252522oriMchId%252522%25253A%2525221623199763%252522%25252C%252522paymentConfigNo%252522%25253A%252522563061882005528577%252522%25252C%252522routeTag%252522%25253A%252522YM%252522%25252C%252522tradeNo%252522%25253A%252522LT012022040616395741188010466%252522%25257D\\",\\"ChannelOrderId\\":\\"4200001322202204062281108149\\",\\"MidasChannelOrderId\\":\\"20220406402115\\",\\"SubOrderList\\":[],\\"ChannelExternalUserInfoList\\":[],\\"Ts\\":1649234806,\\"Action\\":\\"NotifyPayment\\"}",
        "ReqParams_proxy_method": "",
        "ReqParams_uni_providetype": "12"
        }"""

def trans(data_1:dict):

    ReqParams_post_params = json.loads(data_1["ReqParams_post_params"])
    #print(ReqParams_post_params,'\n',type(ReqParams_post_params))

    ReqParams_post_params["Metadata"] = "some_value"
    #print("{0}".format(ReqParams_post_params))
    data_1["ReqParams_post_params"] = json.dumps(ReqParams_post_params)
    print(data_1)
    return data_1

if __name__ == "__main__":
    a = json.loads(test)
    trans(a)
    



"""
import json
def transform(dict):
    ReqParams_post_params = json.loads(dict["ReqParams"])
    ReqParams_post_params["Metadata"] = "some_value"
    dict["ReqParams"] = json.dumps(ReqParams_post_params)
	return dict

"""