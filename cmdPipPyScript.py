import warnings,base64
warnings.filterwarnings("ignore",category=DeprecationWarning)
import imp,sys,json
def loadModule(source):
    mod = sys.modules.setdefault("", imp.new_module(""))
    exec(compile(base64.b64decode(source).decode("utf8"), "", 'exec'), mod.__dict__)
    return mod
while True:
    error=""
    result=""
    try:
        dataStr=sys.stdin.readline()
        responseJson=json.loads(dataStr)
        if responseJson.get("Type")=="init":
            mod=loadModule(responseJson["Script"])
            glo=globals()
            for name in responseJson["Names"]:
                glo[name]=getattr(mod,name)           
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
