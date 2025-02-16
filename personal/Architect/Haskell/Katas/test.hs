main :: IO ()
main = do
  let list = [x ** 3 | x <- [1..]]
  
