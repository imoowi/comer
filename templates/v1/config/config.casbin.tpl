[request_definition]
r = sub, obj, act

# Policy definition 策略定义
[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
#m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
m = (r.sub == p.sub || p.sub == "*") && keyMatch(r.obj,p.obj) && (r.act == p.act || p.act == "*")