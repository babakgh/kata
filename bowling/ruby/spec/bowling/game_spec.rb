# frozen_string_literal: true

require 'rspec'
require_relative '../../lib/bowling/game'

describe Bowling::Game do
  let(:game) { described_class.new }

  def roll_once(pins)
    game.roll(pins)
  end

  def roll_frame(frame)
    game.roll(frame[0])
    game.roll(frame[1])
  end

  def roll_many(numbers, pins)
    numbers.times.each { |_| game.roll(pins) }
  end

  def roll_spare
    roll_many(2, 5)
  end

  def roll_spare_and_extra_ball(bonus_pins)
    roll_spare
    roll_once(bonus_pins)
  end

  def roll_strike
    roll_once(10)
  end

  context 'when a gutter game is played,' do
    it 'results in a score of zero' do
      roll_many(20, 0)
      expect(game.score).to eq(0)
    end
  end

  context 'when all rolls are 1s,' do
    it 'results in a score of 20' do
      roll_many(20, 1)
      expect(game.score).to eq(20)
    end
  end

  context 'when a spare is rolled,' do
    it 'applies the next roll as a bonus to the score' do
      roll_spare
      roll_once(3)
      roll_many(17, 0)

      expect(game.score).to eq(16)
    end
  end

  context 'when a strike is rolled,' do
    it 'applies the next two rolls as a bonus to the score' do
      roll_strike
      roll_frame([3, 4])
      roll_many(16, 0)

      expect(game.score).to eq(24)
    end
  end

  context 'when a perfect game is played,' do
    it 'results in a score of 300' do
      roll_many(12, 10)
      expect(game.score).to eq(300)
    end
  end

  context 'when a game is played,', ignore: true do
    it 'prints result' do
      roll_strike
      roll_frame([3, 4])
      roll_many(14, 0)
      roll_spare_and_extra_ball(5)

      expect(game.print_score_sheet).to eq('17#24#24#24#24#24#24#24#24#39')
    end
  end
end
