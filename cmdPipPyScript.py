import warnings,base64,sys,json
warnings.filterwarnings("ignore",category=DeprecationWarning)
while True:
    dataStr=sys.stdin.readline()
    error=""
    result=""
    try:
        responseJson=json.loads(dataStr)
        if responseJson.get("Type")=="exec":
            exec(base64.b64decode(responseJson["Script"]).decode("utf8"),globals(),locals())
        elif responseJson.get("Type")=="eval":
            result=eval(base64.b64decode(responseJson["Script"]).decode("utf8"),globals(),locals())
        elif responseJson.get("Type")=="init":
            for modulePath in responseJson["ModulePath"]:
                sys.path.append(modulePath)
        elif responseJson.get("Type")=="call":
            if responseJson.get("Args"):
                result=globals()[responseJson.get("Func")](*responseJson["Args"])
            else:
                result=globals()[responseJson.get("Func")]()
        else:
            error="未知的类型"
            result=dataStr
    except Exception as e:
        error=str(e)
        result=dataStr
    sys.stdout.write("##gospider@start##"+json.dumps({"Result":result,"Error":error})+"##gospider@end##")
