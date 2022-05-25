function PastHandler()
    local mq = require("gix-mq")
    while true do
        local val = mq.getPayLoad("issue")
        if (val == nil) then
            break
        end
        print(val)
        PushMySQL(val)
    end
end

function PushMySQL(context)
    local db = require("lua-mysql")
    local mysql = db.connect("root", "root", "127.0.0.1", "3306", "test")
    db.ping(mysql)
    db.insert(mysql, string.format("INSERT INTO demo(name, id) VALUES('%s', '1')", context))
end
