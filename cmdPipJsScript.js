_require = require
_arguments = arguments

process.stdin.on('data', (data) => {
    let result
    let error
    try{
        dataStr=data.toString()
        let responseJson=JSON.parse(dataStr)
        if (responseJson.Type=="exec"){
            global.eval(Buffer.from(responseJson.Script, 'base64').toString('utf-8'))
        }else if (responseJson.Type=="eval"){
            result=global.eval(Buffer.from(responseJson.Script, 'base64').toString('utf-8'))
        }else if (responseJson.Type=="init"){
            responseJson.ModulePath.forEach((path)=>{module.paths.push(path)})        
        }else if (responseJson.Type=="call"){
            if (responseJson.Args){
                result=global.eval(`${responseJson.Func}`)(...responseJson.Args)
            }else{
                result=global.eval(`${responseJson.Func}`)()
            }
        }else{
            error="未知的类型"
            result=dataStr
        }
    }catch(e){
        error=e.stack
        result=dataStr
    }
    process.stdout.write("##gospider@start##"+JSON.stringify({
        Result:result,
        Error:error,
    })+"##gospider@end##")
});