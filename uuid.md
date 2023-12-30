# url

-- http://196.221.166.82:5002/api/receipts/uuid

# body example

-- {"serials" : "311359","store" : 1}

# response example

-- 32500905205fa9a3f849c6833dbe75f75244b712d9d3221b969bd25ac992fe00

# expected behaviour

-- save the generated uuid on the db and also saves the receipt request body at request_body column on StkTrEInvoideHead
