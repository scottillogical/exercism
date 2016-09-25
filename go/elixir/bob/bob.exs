defmodule Bob do
  def hey(input) do
    whoa = ["GO", "WATCH OUT!"]
    cond do
      input == "WATCH OUT!" -> "Whoa, chill out!"
      Regex.match?(~r{cryogenic}, input) -> "Sure."
      Regex.match?(~r{#{whoa}}, input) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end
end
