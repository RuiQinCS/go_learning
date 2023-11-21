local bizs = KEYS
local res = {}
for i = 1, #bizs do
    local bizHotList = redis.call("zrange", bizs[i],"0" ,"99", "rev", "withscores")
    -- 添加新元素到数组末尾
    res[#res + 1] = bizHotList
end

return res
