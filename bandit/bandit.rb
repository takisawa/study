

class Arm
  def initialize(probability)
    @probability = probability
  end

  def run
    if @probability >= rand
      1
    else
      0
    end
  end
end


[0.8, 0.5, 0.2].each do |prob|
  puts "probability: #{prob}"
  arm = Arm.new(prob)
  5.times do
    p arm.run
  end
end
