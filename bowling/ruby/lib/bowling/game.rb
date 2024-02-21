# frozen_string_literal: true

module Bowling
  class Game
    def initialize
      @rolls = []
    end

    def roll(pins)
      @rolls << pins
    end

    def score
      score_in_details.sum
    end

    def print_score_sheet
      score_sheet.join('#')
    end

    private

    def score_in_details
      detailed_score = []
      10.times.reduce(0) do |first_in_frame, _frame|
        detailed_score << calculate_score_in_frame(first_in_frame)

        first_in_frame + balls_in_frame(first_in_frame)
      end
      detailed_score
    end

    def calculate_score_in_frame(first_in_frame)
      if strike?(first_in_frame)
        10 + next_two_balls_for_strike(first_in_frame)
      elsif spare?(first_in_frame)
        10 + next_ball_for_spare(first_in_frame)
      else
        next_two_balls_in_frame(first_in_frame)
      end
    end

    def score_sheet
      accumulated_score = 0
      score_in_details.map do |score_in_frame|
        accumulated_score += score_in_frame
      end
    end

    def balls_in_frame(first_in_frame)
      strike?(first_in_frame) ? 1 : 2
    end

    def strike?(first_in_frame)
      @rolls[first_in_frame] == 10
    end

    def spare?(first_in_frame)
      @rolls[first_in_frame] + @rolls[first_in_frame + 1] == 10
    end

    def next_two_balls_for_strike(first_in_frame)
      @rolls[first_in_frame + 1] + @rolls[first_in_frame + 2]
    end

    def next_ball_for_spare(first_in_frame)
      @rolls[first_in_frame + 2]
    end

    def next_two_balls_in_frame(first_in_frame)
      @rolls[first_in_frame] + @rolls[first_in_frame + 1]
    end
  end
end
