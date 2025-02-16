import Data.Char (isUpper, toLower, toUpper) --to use capital letters only when checking

shiftList :: Int -> [Char] -> [Char]
shiftList 0 xs = xs
shiftList n xs = drop n xs ++ take n xs

indexing :: [Char] -> [(Char, Char)] -> [Char]
indexing [] _ = []
indexing (a : b) pairs = indexingHelper a pairs : indexing b pairs
  where
    indexingHelper a [] = a
    indexingHelper a ((x, w) : xs)
      | toUpper a == x = if isUpper a then w else toLower w
      | otherwise = indexingHelper a xs


main :: IO ()
main = do
  putStr "The encrypted string is: "
  input <- getLine
  let shift = read input :: Int
  let alp = ['A' .. 'Z']
  let shiftedList = shiftList shift alp
  let combinee = zip alp shiftedList
  
  toCrypt <- getLine

  let crypted = indexing toCrypt combinee
  print crypted


